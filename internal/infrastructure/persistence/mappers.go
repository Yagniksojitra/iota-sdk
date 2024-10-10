package persistence

import (
	"errors"
	stage "github.com/iota-agency/iota-erp/internal/domain/entities/project_stages"
	"github.com/iota-agency/iota-erp/internal/domain/entities/user"
	"time"

	category "github.com/iota-agency/iota-erp/internal/domain/aggregates/expense_category"
	moneyAccount "github.com/iota-agency/iota-erp/internal/domain/aggregates/money_account"
	"github.com/iota-agency/iota-erp/internal/domain/aggregates/order"
	"github.com/iota-agency/iota-erp/internal/domain/aggregates/project"
	"github.com/iota-agency/iota-erp/internal/domain/entities/currency"
	"github.com/iota-agency/iota-erp/internal/domain/entities/payment"
	"github.com/iota-agency/iota-erp/internal/domain/entities/product"
	"github.com/iota-agency/iota-erp/internal/domain/entities/transaction"
	"github.com/iota-agency/iota-erp/internal/domain/entities/unit"
	"github.com/iota-agency/iota-erp/internal/infrastructure/persistence/models"
)

func toDBUnit(unit *unit.Unit) *models.WarehouseUnit {
	return &models.WarehouseUnit{
		ID:        unit.ID,
		Name:      unit.Name,
		CreatedAt: unit.CreatedAt,
		UpdatedAt: unit.UpdatedAt,
	}
}

func toDomainUnit(dbUnit *models.WarehouseUnit) *unit.Unit {
	return &unit.Unit{
		ID:          dbUnit.ID,
		Name:        dbUnit.Name,
		Description: dbUnit.Description,
		CreatedAt:   dbUnit.CreatedAt,
		UpdatedAt:   dbUnit.UpdatedAt,
	}
}

func toDBOrder(data *order.Order) (*models.WarehouseOrder, []*models.OrderItem) {
	dbItems := make([]*models.OrderItem, 0, len(data.Items))
	for _, item := range data.Items {
		dbItems = append(
			dbItems, &models.OrderItem{
				ProductID: item.Product.ID,
				OrderID:   data.ID,
				CreatedAt: data.CreatedAt,
			},
		)
	}
	return &models.WarehouseOrder{
		ID:        data.ID,
		Status:    data.Status.String(),
		Type:      data.Type.String(),
		CreatedAt: data.CreatedAt,
	}, dbItems
}

func toDomainOrder(
	dbOrder *models.WarehouseOrder,
	dbItems []*models.OrderItem,
	dbProduct []*models.WarehouseProduct,
) (*order.Order, error) {
	items := make([]*order.Item, 0, len(dbItems))
	for _, item := range dbItems {
		var orderProduct *models.WarehouseProduct
		for _, p := range dbProduct {
			if p.ID == item.ProductID {
				orderProduct = p
				break
			}
		}
		if orderProduct == nil {
			return nil, errors.New("product not found")
		}
		p, err := toDomainProduct(orderProduct)
		if err != nil {
			return nil, err
		}
		items = append(
			items, &order.Item{
				Product:   p,
				CreatedAt: item.CreatedAt,
			},
		)
	}
	status, err := order.NewStatus(dbOrder.Status)
	if err != nil {
		return nil, err
	}
	typeEnum, err := order.NewType(dbOrder.Type)
	if err != nil {
		return nil, err
	}
	return &order.Order{
		ID:        dbOrder.ID,
		Status:    status,
		Type:      typeEnum,
		CreatedAt: dbOrder.CreatedAt,
		Items:     items,
	}, nil
}

func toDBProduct(entity *product.Product) *models.WarehouseProduct {
	return &models.WarehouseProduct{
		ID:         entity.ID,
		PositionID: entity.PositionID,
		Rfid:       entity.Rfid,
		Status:     entity.Status.String(),
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

func toDomainProduct(dbProduct *models.WarehouseProduct) (*product.Product, error) {
	status, err := product.NewStatus(dbProduct.Status)
	if err != nil {
		return nil, err
	}
	return &product.Product{
		ID:         dbProduct.ID,
		PositionID: dbProduct.PositionID,
		Rfid:       dbProduct.Rfid,
		Status:     status,
		CreatedAt:  dbProduct.CreatedAt,
		UpdatedAt:  dbProduct.UpdatedAt,
	}, nil
}

func toDBTransaction(entity *transaction.Transaction) *models.Transaction {
	return &models.Transaction{
		ID:               entity.ID,
		Amount:           entity.Amount,
		Comment:          entity.Comment,
		AccountingPeriod: entity.AccountingPeriod,
		AmountCurrencyID: entity.AmountCurrencyID,
		TransactionDate:  entity.TransactionDate,
		MoneyAccountID:   entity.MoneyAccountID,
		TransactionType:  entity.TransactionType.String(),
		CreatedAt:        entity.CreatedAt,
	}
}

func toDomainTransaction(dbTransaction *models.Transaction) (*transaction.Transaction, error) {
	_type, err := transaction.NewType(dbTransaction.TransactionType)
	if err != nil {
		return nil, err
	}

	return &transaction.Transaction{
		ID:               dbTransaction.ID,
		Amount:           dbTransaction.Amount,
		TransactionType:  _type,
		Comment:          dbTransaction.Comment,
		AccountingPeriod: dbTransaction.AccountingPeriod,
		TransactionDate:  dbTransaction.TransactionDate,
		MoneyAccountID:   dbTransaction.MoneyAccountID,
		AmountCurrencyID: dbTransaction.AmountCurrencyID,
		CreatedAt:        dbTransaction.CreatedAt,
	}, nil
}

func toDBPayment(entity *payment.Payment) (*models.Payment, *models.Transaction) {
	dbTransaction := &models.Transaction{
		ID:               entity.TransactionID,
		Amount:           entity.Amount,
		Comment:          entity.Comment,
		AccountingPeriod: entity.AccountingPeriod,
		TransactionDate:  entity.TransactionDate,
		MoneyAccountID:   entity.AccountID,
		AmountCurrencyID: entity.CurrencyCode,
		TransactionType:  transaction.Income.String(),
		CreatedAt:        entity.CreatedAt,
	}
	dbPayment := &models.Payment{
		ID:            entity.ID,
		StageID:       entity.StageID,
		TransactionID: entity.TransactionID,
		Transaction:   dbTransaction,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
	return dbPayment, dbTransaction
}

func toDomainPayment(dbPayment *models.Payment) (*payment.Payment, error) {
	if dbPayment.Transaction == nil {
		return nil, errors.New("transaction is nil")
	}
	t, err := toDomainTransaction(dbPayment.Transaction)
	if err != nil {
		return nil, err
	}
	return &payment.Payment{
		ID:               dbPayment.ID,
		Amount:           t.Amount,
		Comment:          t.Comment,
		TransactionDate:  t.TransactionDate,
		AccountingPeriod: t.AccountingPeriod,
		StageID:          dbPayment.StageID,
		User:             &user.User{},
		TransactionID:    dbPayment.TransactionID,
		AccountID:        t.MoneyAccountID,
		CurrencyCode:     t.AmountCurrencyID,
		CreatedAt:        dbPayment.CreatedAt,
		UpdatedAt:        dbPayment.UpdatedAt,
	}, nil
}

func toDBCurrency(entity *currency.Currency) *models.Currency {
	return &models.Currency{
		Code:      entity.Code.String(),
		Name:      entity.Name,
		Symbol:    entity.Symbol.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func toDomainCurrency(dbCurrency *models.Currency) (*currency.Currency, error) {
	code, err := currency.NewCode(dbCurrency.Code)
	if err != nil {
		return nil, err
	}
	symbol, err := currency.NewSymbol(dbCurrency.Symbol)
	if err != nil {
		return nil, err
	}
	return &currency.Currency{
		Code:   code,
		Name:   dbCurrency.Name,
		Symbol: symbol,
	}, nil
}

func toDBExpenseCategory(entity *category.ExpenseCategory) *models.ExpenseCategory {
	return &models.ExpenseCategory{
		ID:               entity.ID,
		Name:             entity.Name,
		Description:      &entity.Description,
		Amount:           entity.Amount,
		AmountCurrencyID: entity.Currency.Code.String(),
		CreatedAt:        entity.CreatedAt,
		UpdatedAt:        entity.UpdatedAt,
	}
}

func toDomainExpenseCategory(dbCategory *models.ExpenseCategory) (*category.ExpenseCategory, error) {
	currencyEntity, err := toDomainCurrency(&dbCategory.AmountCurrency)
	if err != nil {
		return nil, err
	}
	return &category.ExpenseCategory{
		ID:          dbCategory.ID,
		Name:        dbCategory.Name,
		Description: *dbCategory.Description,
		Amount:      dbCategory.Amount,
		Currency:    *currencyEntity,
		CreatedAt:   dbCategory.CreatedAt,
		UpdatedAt:   dbCategory.UpdatedAt,
	}, nil
}

func toDomainProject(dbProject *models.Project) *project.Project {
	return &project.Project{
		Id:          dbProject.ID,
		Name:        dbProject.Name,
		Description: dbProject.Description,
		CreatedAt:   dbProject.CreatedAt,
		UpdatedAt:   dbProject.UpdatedAt,
	}
}

func toDBProject(entity *project.Project) *models.Project {
	return &models.Project{
		ID:          entity.Id,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func toDomainMoneyAccount(dbAccount *models.MoneyAccount) (*moneyAccount.Account, error) {
	currencyEntity, err := toDomainCurrency(dbAccount.Currency)
	if err != nil {
		return nil, err
	}
	return &moneyAccount.Account{
		ID:            dbAccount.ID,
		Name:          dbAccount.Name,
		AccountNumber: dbAccount.AccountNumber,
		Balance:       dbAccount.Balance,
		Currency:      *currencyEntity,
		Description:   dbAccount.Description,
		CreatedAt:     dbAccount.CreatedAt,
		UpdatedAt:     dbAccount.UpdatedAt,
	}, nil
}

func toDBMoneyAccount(entity *moneyAccount.Account) *models.MoneyAccount {
	return &models.MoneyAccount{
		ID:                entity.ID,
		Name:              entity.Name,
		AccountNumber:     entity.AccountNumber,
		Balance:           entity.Balance,
		BalanceCurrencyID: entity.Currency.Code.String(),
		Currency:          toDBCurrency(&entity.Currency),
		Description:       entity.Description,
		CreatedAt:         entity.CreatedAt,
		UpdatedAt:         entity.UpdatedAt,
	}
}

func toDomainProjectStage(dbStage *models.ProjectStage) *stage.ProjectStage {
	return &stage.ProjectStage{
		Id:        dbStage.ID,
		Name:      dbStage.Name,
		CreatedAt: dbStage.CreatedAt,
		UpdatedAt: dbStage.UpdatedAt,
	}
}

func toDBProjectStage(entity *stage.ProjectStage) *models.ProjectStage {
	return &models.ProjectStage{
		ID:        entity.Id,
		Name:      entity.Name,
		ProjectID: entity.ProjectID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

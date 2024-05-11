package postgres

import (
	"fmt"
	"module/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderPG struct {
	db *sqlx.DB
}

func NewOrderPG(db *sqlx.DB) *OrderPG {
	return &OrderPG{db: db}
}

func (r *OrderPG) CreateOrder(o models.OrderT) (string, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return "", err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var order_uid string
	query1 := fmt.Sprintf(
		`insert into %s (
			order_uid, 
			track_number, 
			entry, 
			locale, 
			internal_signature, 
			customer_id, 
			delivery_service, 
			shardkey, 
			sm_id, 
			date_created, 
			oof_shard
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11)
		returning order_uid`, OrdersTable)
	row := tx.QueryRow(query1, o.OrderUid, o.TrackNumber, o.Entry, o.Locale,
		o.InternalSignature, o.CustomerId, o.DeliveryService, o.Shardkey,
		o.SmId, o.DateCreated, o.OofShard)
	err = row.Scan(&order_uid)
	if err != nil {
		return "", err
	}

	query2 := fmt.Sprintf(
		`insert into %s (
			order_uid,
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11)
		returning order_uid`, PaymentsTable)
	row = tx.QueryRow(query2, o.OrderUid, o.Payment.Transaction,
		o.Payment.RequestId, o.Payment.Currency, o.Payment.Provider,
		o.Payment.Amount, o.Payment.PaymentDt, o.Payment.Bank,
		o.Payment.DeliveryCost, o.Payment.GoodsTotal, o.Payment.CustomFee)
	err = row.Scan(&order_uid)
	if err != nil {
		return "", err
	}

	query3 := fmt.Sprintf(
		`insert into %s (
			order_uid,
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11,
			$12)
		returning item_id`, ItemsTable)
	for _, item := range o.Items {
		var id int
		row = tx.QueryRow(query3, o.OrderUid, item.ChrtId, item.TrackNumber, item.Price,
			item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId,
			item.Brand, item.Status)
		if err := row.Scan(&id); err != nil {
			return "", err
		}
	}

	query4 := fmt.Sprintf(
		`insert into %s (
			order_uid,
			name,
			phone,
			zip,
			city,
			address,
			region,
			email
		) values (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8)
		returning order_uid`, DeliveriesTable)
	row = tx.QueryRow(query4, o.OrderUid, o.Delivery.Name, o.Delivery.Phone,
		o.Delivery.Zip, o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email)
	err = row.Scan(&order_uid)
	if err != nil {
		return "", err
	}
	return order_uid, nil
}

func (r *OrderPG) GetOrders() ([]models.OrderT, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var orders []models.OrderT
	query1 := fmt.Sprintf(`select * from %s`, OrdersTable)
	err = tx.Select(&orders, query1)
	if err != nil {
		return nil, err
	}
	for i := range orders {
		var tmpPayment models.PaymentT
		query2 := fmt.Sprintf(
			`select
				transaction,
				request_id,
				currency,
				provider,
				amount,
				payment_dt,
				bank,
				delivery_cost,
				goods_total,
				custom_fee 
			from %s 
			where order_uid=$1`, PaymentsTable)
		err = tx.Get(&tmpPayment, query2, orders[i].OrderUid)
		if err != nil {
			return nil, err
		}
		orders[i].Payment = tmpPayment

		var tmpItems []models.ItemT
		query3 := fmt.Sprintf(
			`select 
				chrt_id,
				track_number,
				price,
				rid,
				name,
				sale,
				size,
				total_price,
				nm_id,
				brand,
				status 
			from %s 
			where order_uid=$1`, ItemsTable)
		err = tx.Select(&tmpItems, query3, orders[i].OrderUid)
		if err != nil {
			return nil, err
		}
		orders[i].Items = tmpItems

		var tmpDelivery models.DeliveryT
		query4 := fmt.Sprintf(
			`select 
				name,
				phone,
				zip,
				city,
				address,
				region,
				email
			from %s 
			where order_uid=$1`, DeliveriesTable)
		err = tx.Get(&tmpDelivery, query4, orders[i].OrderUid)
		if err != nil {
			return nil, err
		}
		orders[i].Delivery = tmpDelivery
	}

	return orders, nil
}

func (r *OrderPG) GetOrderByUid(uid string) (models.OrderT, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return models.OrderT{}, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var order models.OrderT
	query1 := fmt.Sprintf(`select * from %s where order_uid=$1`, OrdersTable)
	err = tx.Get(&order, query1, uid)
	if err != nil {
		return models.OrderT{}, err
	}
	var tmpPayment models.PaymentT
	query2 := fmt.Sprintf(
		`select
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee 
		from %s 
		where order_uid=$1`, PaymentsTable)
	err = tx.Get(&tmpPayment, query2, uid)
	if err != nil {
		return models.OrderT{}, err
	}
	order.Payment = tmpPayment

	var tmpItems []models.ItemT
	query3 := fmt.Sprintf(
		`select 
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status 
		from %s 
		where order_uid=$1`, ItemsTable)
	err = tx.Select(&tmpItems, query3, uid)
	if err != nil {
		return models.OrderT{}, err
	}
	order.Items = tmpItems

	var tmpDelivery models.DeliveryT
	query4 := fmt.Sprintf(
		`select 
			name,
			phone,
			zip,
			city,
			address,
			region,
			email
		from %s 
		where order_uid=$1`, DeliveriesTable)
	err = tx.Get(&tmpDelivery, query4, uid)
	if err != nil {
		return models.OrderT{}, err
	}
	order.Delivery = tmpDelivery

	return order, nil
}

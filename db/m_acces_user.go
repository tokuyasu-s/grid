package db

import "time"

type m_acces_user struct {
	m_acc_id           int
	m_user_id          int
	m_user_name        string
	m_user_age         int
	m_user_birthday    string
	m_user_gender      string
	m_user_status      string
	m_user_regist_date time.Time
	m_user_update_date time.Time
	m_user_bikou       string
	m_user_stay        string
}

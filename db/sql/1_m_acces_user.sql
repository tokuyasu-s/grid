create table m_acces_user(
 m_acc_id bigserial not null primary key,
 m_user_id bigint not null,
 m_user_name text not null,
 m_user_age int not null,
 m_user_birthday date null,
 m_user_gender varchar(1) not null,
 m_user_status varchar(1) not null default '1',
 m_user_regist_date timestamp not null,
 m_user_update_date timestamp not null,
 m_user_bikou text,
 m_user_stay text
)
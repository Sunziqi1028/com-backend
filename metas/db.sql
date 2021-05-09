create table if not exists comer_tbl(
    id bigint not null auto_increament,
    uin bigint not null comment 'comunion comer unique identifier',
    address varchar(50) not null default '' comment 'comunion comer could save some useful info on block chain with this address',
    comer_id varchar(35) not null default '' comment 'comunion comer UUID',
    nick varchar(50) not null default '' comment 'comunion comer nick name',
    avatar varchar(255) not null default '' comment 'comunion avatar link address',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id),
    index comer_uin_idx(uin)
);

create table if not exists account_tbl(
    id bigint not null auto_increament,
    uin bigint not null comment 'comunion comer unique identifier',
    oin varchar(100) not null comment 'comunion comer outer account unique identifier, wallet will be public key and Oauth is the OauthID',
    main smallint not null default 0 comment 'comunion comer use this account as main account',
    nick varchar(50) not null default '' comment 'comunion comer nick name',
    avatar varchar(255) not null default '' comment 'comunion avatar link address',
    category int not null default 0 comment 'comunion outer account type 1 for wallet 2 for Oauth',
    type int not null default 0 comment '1 for github 2 for metamask 3 for twitter 4 for facebook 5 for likedin 6 for iamtoken',
    linked smallint not null default 0 comment '0 for unlink 1 for linked',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id),
    index comer_uin_idx(uin),
    index comer_oin_uin_idx(oin, uin)
);


create table if not exists bounty_tbl(
    id bigint not null auto_increament,
    ifentifier bigint not null comment 'comunion bounty identifier',
    startup_identifier bigint not null comment 'comunion startup identifier',
    name varchar(50) not null comment 'comunion bounty name',
    description varchar(255) not null comment 'comunion shorter description of this bounty',
    created_by bigint not null comment 'comunion bounty creator use comunion uin',
    address varchar(50) not null comment 'comunion bounty block chain address',
    introduce text not null comment 'comunion bounty introduce',
    state int not null default 0 comment 'comunion bounty state 1 for open 2 for processing 3 for closed',
    start_at datetime default null comment 'when this bounty start work',
    closed_at datetime default null comment 'when this bounty end work and closed',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id),
    index bounty_identifier_idx(identifier) using btree,
    index bounty_startup_idx(startup_identifier) using btree,
    index bounty_creator_idx(created_by) using btree,
);

create table if not exists bounty_comer_rel_tbl(
    id bigint not null auto_increament,
    bounty_identifier bigint not null comment 'comunion bounty identifier',
    comer_uin bigint not null comment 'comunion comer UIN',
    state smallint not null default 0 comment 'comer active state 1 for submit 2 for start work 3 for end work',
    type smallint not null default 0 comment 'comer 1 for participation 2 for following',
    deleted smallint not null default 0 comment 'if the bind releation is deleted?',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id),
    index bounty_identifier_comer_idx(bounty_identifier, comer_uin, state, type) using btree
);


create table if not exists comer_profile_tbl(
    id bigint not null auto_increament,
    uin bigint not null default 0 comment 'comunion comer uin',
    remark varchar(30) not null default '' comment 'comunion profile name',
    identifier bigint not null comment 'comunion comer profile',
    name varchar(50) not null comment 'comunion comer which name comer wanna displaying to other',
    description varchar(255) not null comment 'comunion comer profile description',
    email varchar(100) not null comment 'comunion comer profile email',
    version int not null default 1 comment 'comunion comer profile version',
    skills varchar(30) not null default '' comment 'comunion comer skills list split by comma',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id),
    index profile_identifier_idx(identifier) using btree,
    index comer_uin_index(uin) using btree
);

create table if not exists comer_profile_skill_tag_tbl(
    id bigint not null auto_increament,
    name varchar(50) not null comment 'comer skill tag name',
    valid smallint not null default 1 comment 'if this skill tag is avlidable',
    create_at datetime not null default current_timestamp,
    update_at datetime not null default current_timestamp on update current_timestamp,
    primary key(id)
);



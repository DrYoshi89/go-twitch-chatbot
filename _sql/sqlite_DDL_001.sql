--
-- SQLITE
--

CREATE TABLE IF NOT EXISTS sys (
	tx_sys INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.sys

	id_autostart INTEGER NOT NULL CHECK (id_autostart IN (0, 1)), -- Inactive, Active
);

CREATE TABLE IF NOT EXISTS user (
	cd_user INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.user
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1, 2)) DEFAULT 1, -- Active, Suspended, Banned
	id_permission INTEGER NOT NULL CHECK (id_permission IN (0, 1, 2)) DEFAULT 2, -- Admin, Mod, Human
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	cd_user_group INTEGER NOT NULL,
	tx_auth TEXT,
	cd_ban INTEGER,
	tx_nick TEXT NOT NULL,
	bl_avatar BLOB,

	tx_name TEXT,
	tx_description TEXT,

	FOREIGN KEY (cd_user_group) REFERENCES user_group (cd_user_group) ON DELETE SET NULL ON UPDATE SET NULL,
	FOREIGN KEY (cd_ban) REFERENCES ban_type (cd_ban) ON DELETE SET NULL ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS user_group (
	cd_user_group INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.user_group
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1, 2)) DEFAULT 1, -- Active, Suspended, Banned
	id_permission INTEGER NOT NULL CHECK (id_permission IN (0, 1, 2)) DEFAULT 2, -- Admin, Mod, Human
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_name TEXT,
	tx_description TEXT,

	bl_avatar BLOB
);

CREATE TABLE IF NOT EXISTS user_settings (
	cd_user_settings INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.user_settings
	cd_user INTEGER NOT NULL,

	id_darkmode INTEGER CHECK (id_darkmode IN (0, 1)), -- Inactive, Active

	FOREIGN KEY (cd_user) REFERENCES user (cd_user) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS channel (
	cd_channel INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.channel
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	in_autojoin INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- False, Tue
	tx_channel TEXT NOT NULL,
	bl_thumb BLOB

);

CREATE TABLE IF NOT EXISTS command (
	cd_command INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.command
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	cd_channel INTEGER NOT NULL,

	tx_name TEXT NOT NULL,

	FOREIGN KEY (cd_channel) REFERENCES channel (cd_channel) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS platform (
	cd_platform INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.platform
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active

	tx_auth TEXT NOT NULL,
	tx_url_auth TEXT NOT NULL,
	tx_url_site TEXT,

	tx_name TEXT NOT NULL,
	tx_description TEXT,
	bl_thumb BLOB
);

CREATE TABLE IF NOT EXISTS ban_list (
	cd_ban_list INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.ban_list
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	cd_ban_type INTEGER DEFAULT NULL,

	tx_description TEXT NOT NULL,

	FOREIGN KEY (cd_ban_type) REFERENCES ban_type (cd_ban_type) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS ban_type (
	cd_ban_type INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.ban_type
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_description TEXT
);

CREATE TABLE IF NOT EXISTS message (
	cd_message INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.message
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Inactive, Active
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_channel TEXT NOT NULL,
	tx_raw TEXT NOT NULL,
	tx_message TEXT NOT NULL,
	tx_user TEXT NOT NULL,

	cd_user TEXT,

	FOREIGN KEY (cd_user) REFERENCES users (cd_user) ON DELETE NO ACTION ON UPDATE NO ACTION
);

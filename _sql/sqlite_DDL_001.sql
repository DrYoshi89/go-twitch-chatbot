--
-- SQLITE
--
CREATE TABLE IF NOT EXISTS user (
	cd_user INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.user
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1, 2)), -- Active, Suspended, Banned
	id_permission INTEGER NOT NULL CHECK (id_permission IN (0, 1, 2, 3)), -- admin, mod, sub, registred
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_auth TEXT,
	cd_ban INTEGER,
	tx_nick TEXT NOT NULL,
	bl_avatar BLOB,

	tx_name TEXT,
	tx_description TEXT,

	FOREIGN KEY (cd_ban) REFERENCES ban_type (cd_ban) ON DELETE SET NULL ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS user_settings (
	cd_user_settings INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.user_settings
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1, 2)), -- Active, Suspended, Banned
	id_permission INTEGER NOT NULL CHECK (id_permission IN (0, 1, 2, 3)), -- admin, mod, sub, registred
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	cd_user INTEGER NOT NULL,

	id_darkmode INTEGER CHECK (id_darkmode IN (0, 1)), -- False, True

	FOREIGN KEY (cd_user) REFERENCES user (cd_user) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS platform (
	cd_platform INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.platform
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Active, Inactive

	tx_auth TEXT NOT NULL,
	tx_url_auth TEXT NOT NULL,
	tx_url_site TEXT,

	tx_name TEXT NOT NULL,
	tx_description TEXT,

	FOREIGN KEY (cd_user) REFERENCES user (cd_user) ON DELETE CASCADE ON UPDATE CASCADE,
);

CREATE TABLE IF NOT EXISTS ban_list (
	cd_ban_list INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.ban_list
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Active, Inactive
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	cd_ban_type INTEGER DEFAULT NULL,

	tx_description TEXT NOT NULL,

	FOREIGN KEY (cd_ban_type) REFERENCES ban_type (cd_ban_type) ON DELETE CASCADE ON UPDATE CASCADE,
);


CREATE TABLE IF NOT EXISTS ban_type (
	cd_ban_type INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.ban_type
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Active, Inactive
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_description TEXT,
);

CREATE TABLE IF NOT EXISTS message (
	cd_message INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.message
	id_status INTEGER NOT NULL CHECK (id_status IN (0, 1)) DEFAULT 1, -- Active, Inactive
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	tx_channel TEXT NOT NULL,
	tx_raw TEXT NOT NULL,
	tx_message TEXT NOT NULL,
	tx_user TEXT NOT NULL,

	cd_user TEXT,

	FOREIGN KEY (cd_user) REFERENCES users (cd_user) ON DELETE NO ACTION ON UPDATE NO ACTION
);

--
-- SQLITE
--

-- add admin
INSERT INTO user (id_status, id_permission, dt_created, tx_auth, cd_ban, tx_nick, bl_avatar, tx_name, tx_description)
	VALUES (1, 0, datetime(), '',  null, 'admin', null, null, null);

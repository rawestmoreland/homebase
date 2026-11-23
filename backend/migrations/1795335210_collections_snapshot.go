package migrations

import (
	"ganigeorgiev.com/pocketbase/pocketbase/core"
	m "ganigeorgiev.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `[
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1582905952",
						"max": 0,
						"min": 0,
						"name": "method",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_2279338944",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_mfas_collectionRef_recordRef` + "`" + ` ON ` + "`" + `_mfas` + "`" + ` (collectionRef,recordRef)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_mfas",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 8,
						"help": "",
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 0,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": true,
						"id": "text3866985172",
						"max": 0,
						"min": 0,
						"name": "sentTo",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_1638494021",
				"indexes": [
					"CREATE INDEX ` + "`" + `idx_otps_collectionRef_recordRef` + "`" + ` ON ` + "`" + `_otps` + "`" + ` (collectionRef, recordRef)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_otps",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text2462348188",
						"max": 0,
						"min": 0,
						"name": "provider",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1044722854",
						"max": 0,
						"min": 0,
						"name": "providerId",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_2281828961",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_externalAuths_record_provider` + "`" + ` ON ` + "`" + `_externalAuths` + "`" + ` (collectionRef, recordRef, provider)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_externalAuths_collection_provider` + "`" + ` ON ` + "`" + `_externalAuths` + "`" + ` (collectionRef, provider, providerId)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_externalAuths",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"createRule": null,
				"deleteRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text455797646",
						"max": 0,
						"min": 0,
						"name": "collectionRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text127846527",
						"max": 0,
						"min": 0,
						"name": "recordRef",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text4228609354",
						"max": 0,
						"min": 0,
						"name": "fingerprint",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"id": "pbc_4275539003",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_authOrigins_unique_pairs` + "`" + ` ON ` + "`" + `_authOrigins` + "`" + ` (collectionRef, recordRef, fingerprint)"
				],
				"listRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId",
				"name": "_authOrigins",
				"system": true,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != '' && recordRef = @request.auth.id && collectionRef = @request.auth.collectionId"
			},
			{
				"authAlert": {
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>We noticed a login to your {APP_NAME} account from a new location:</p>\n<p><em>{ALERT_INFO}</em></p>\n<p><strong>If this wasn't you, you should immediately change your {APP_NAME} account password to revoke access from all other locations.</strong></p>\n<p>If this was you, you may disregard this email.</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "Login from a new location"
					},
					"enabled": true
				},
				"authRule": "",
				"authToken": {
					"duration": 86400
				},
				"confirmEmailChangeTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Confirm your {APP_NAME} new email address"
				},
				"createRule": null,
				"deleteRule": null,
				"emailChangeToken": {
					"duration": 1800
				},
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 0,
						"help": "",
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 8,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "[a-zA-Z0-9]{50}",
						"help": "",
						"hidden": true,
						"id": "text2504183744",
						"max": 60,
						"min": 30,
						"name": "tokenKey",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"exceptDomains": [],
						"help": "",
						"hidden": false,
						"id": "email3885137012",
						"name": "email",
						"onlyDomains": [],
						"presentable": false,
						"required": true,
						"system": true,
						"type": "email"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool1547992806",
						"name": "emailVisibility",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool256245529",
						"name": "verified",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": true,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": true,
						"type": "autodate"
					}
				],
				"fileToken": {
					"duration": 180
				},
				"id": "pbc_3142635823",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey_pbc_3142635823` + "`" + ` ON ` + "`" + `_superusers` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_email_pbc_3142635823` + "`" + ` ON ` + "`" + `_superusers` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''"
				],
				"listRule": null,
				"manageRule": null,
				"mfa": {
					"duration": 600,
					"enabled": false,
					"rule": ""
				},
				"name": "_superusers",
				"oauth2": {
					"enabled": false,
					"mappedFields": {
						"avatarURL": "",
						"id": "",
						"name": "",
						"username": ""
					}
				},
				"otp": {
					"duration": 180,
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>Your one-time password is: <strong>{OTP}</strong></p>\n<p><i>If you didn't ask for the one-time password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "OTP for {APP_NAME}"
					},
					"enabled": false,
					"length": 8
				},
				"passwordAuth": {
					"enabled": true,
					"identityFields": [
						"email"
					]
				},
				"passwordResetToken": {
					"duration": 1800
				},
				"resetPasswordTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-password-reset/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Reset your {APP_NAME} password"
				},
				"system": true,
				"type": "auth",
				"updateRule": null,
				"verificationTemplate": {
					"body": "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-verification/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p><i>If you didn't recently register, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Verify your {APP_NAME} email"
				},
				"verificationToken": {
					"duration": 86400
				},
				"viewRule": null
			},
			{
				"authAlert": {
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>We noticed a login to your {APP_NAME} account from a new location:</p>\n<p><em>{ALERT_INFO}</em></p>\n<p><strong>If this wasn't you, you should immediately change your {APP_NAME} account password to revoke access from all other locations.</strong></p>\n<p>If this was you, you may disregard this email.</p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "Login from a new location"
					},
					"enabled": true
				},
				"authRule": "",
				"authToken": {
					"duration": 432000
				},
				"confirmEmailChangeTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to confirm your new email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-email-change/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Confirm new email</a>\n</p>\n<p><i>If you didn't ask to change your email address, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Confirm your {APP_NAME} new email address"
				},
				"createRule": "",
				"deleteRule": "id = @request.auth.id",
				"emailChangeToken": {
					"duration": 1800
				},
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cost": 0,
						"help": "",
						"hidden": true,
						"id": "password901924565",
						"max": 0,
						"min": 8,
						"name": "password",
						"pattern": "",
						"presentable": false,
						"required": true,
						"system": true,
						"type": "password"
					},
					{
						"autogeneratePattern": "[a-zA-Z0-9]{50}",
						"help": "",
						"hidden": true,
						"id": "text2504183744",
						"max": 60,
						"min": 30,
						"name": "tokenKey",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"exceptDomains": [],
						"help": "",
						"hidden": false,
						"id": "email3885137012",
						"name": "email",
						"onlyDomains": [],
						"presentable": false,
						"required": true,
						"system": true,
						"type": "email"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool1547992806",
						"name": "emailVisibility",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool256245529",
						"name": "verified",
						"presentable": false,
						"required": false,
						"system": true,
						"type": "bool"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1579384326",
						"max": 255,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "file376926767",
						"maxSelect": 1,
						"maxSize": 0,
						"mimeTypes": [
							"image/jpeg",
							"image/png",
							"image/svg+xml",
							"image/gif",
							"image/webp"
						],
						"name": "avatar",
						"presentable": false,
						"protected": false,
						"required": false,
						"system": false,
						"thumbs": [],
						"type": "file"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"fileToken": {
					"duration": 180
				},
				"id": "_pb_users_auth_",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_email__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''"
				],
				"listRule": "id = @request.auth.id",
				"manageRule": null,
				"mfa": {
					"duration": 600,
					"enabled": false,
					"rule": ""
				},
				"name": "users",
				"oauth2": {
					"enabled": false,
					"mappedFields": {
						"avatarURL": "avatar",
						"id": "",
						"name": "name",
						"username": ""
					}
				},
				"otp": {
					"duration": 180,
					"emailTemplate": {
						"body": "<p>Hello,</p>\n<p>Your one-time password is: <strong>{OTP}</strong></p>\n<p><i>If you didn't ask for the one-time password, you can ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
						"subject": "OTP for {APP_NAME}"
					},
					"enabled": false,
					"length": 8
				},
				"passwordAuth": {
					"enabled": true,
					"identityFields": [
						"email"
					]
				},
				"passwordResetToken": {
					"duration": 1800
				},
				"resetPasswordTemplate": {
					"body": "<p>Hello,</p>\n<p>Click on the button below to reset your password.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-password-reset/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Reset password</a>\n</p>\n<p><i>If you didn't ask to reset your password, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Reset your {APP_NAME} password"
				},
				"system": false,
				"type": "auth",
				"updateRule": "id = @request.auth.id",
				"verificationTemplate": {
					"body": "<p>Hello,</p>\n<p>Thank you for joining us at {APP_NAME}.</p>\n<p>Click on the button below to verify your email address.</p>\n<p>\n  <a class=\"btn\" href=\"{APP_URL}/_/#/auth/confirm-verification/{TOKEN}\" target=\"_blank\" rel=\"noopener\">Verify</a>\n</p>\n<p><i>If you didn't recently register, please ignore this email.</i></p>\n<p>\n  Thanks,<br/>\n  {APP_NAME} team\n</p>",
					"subject": "Verify your {APP_NAME} email"
				},
				"verificationToken": {
					"duration": 86400
				},
				"viewRule": "id = @request.auth.id"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1579384326",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text2560465762",
						"max": 0,
						"min": 0,
						"name": "slug",
						"pattern": "^[a-z0-9-]+$",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "IANA name, e.g. America/Chicago",
						"hidden": false,
						"id": "text922858135",
						"max": 0,
						"min": 0,
						"name": "timezone",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "usd, cad, eur, ...",
						"hidden": false,
						"id": "text1767278655",
						"max": 0,
						"min": 0,
						"name": "currency",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text3764779314",
						"max": 0,
						"min": 0,
						"name": "stripe_account_id",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select3545045024",
						"maxSelect": 0,
						"name": "stripe_status",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"none",
							"pending",
							"active",
							"restricted"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "json3846545605",
						"maxSize": 0,
						"name": "settings",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "json"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3736699063",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_ygnu5jfedh` + "`" + ` ON ` + "`" + `associations` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.association ?= id &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.status ?= \"active\"",
				"name": "associations",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= id &&\n@collection.memberships:mine.role ?= \"admin\" &&\n@collection.memberships:mine.status ?= \"active\" &&\n@request.body.stripe_account_id:changed = false && @request.body.stripe_status:changed = false &&\n@request.body.slug:changed = false &&\n@request.body.currency:changed = false\n",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.association ?= id &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.status ?= \"active\""
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\"",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation2375276105",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "user",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": true,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select1466534506",
						"maxSelect": 0,
						"name": "role",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "select",
						"values": [
							"admin",
							"board",
							"resident"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "select2063623452",
						"maxSelect": 0,
						"name": "status",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "select",
						"values": [
							"invited",
							"active",
							"disabled"
						]
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_1990943746",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_0ek8s12uwb` + "`" + ` ON ` + "`" + `memberships` + "`" + ` (\n  ` + "`" + `association` + "`" + `,\n  ` + "`" + `user` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\" && (@request.auth.id = user || (\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")\n))",
				"name": "memberships",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\" &&\n@request.body.user:changed = false &&\n@request.body.association:changed = false",
				"viewRule": "@request.auth.id != \"\" && (@request.auth.id = user || (\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")\n))"
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\"",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text223244161",
						"max": 0,
						"min": 0,
						"name": "address",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text3703245907",
						"max": 0,
						"min": 0,
						"name": "unit",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text193014043",
						"max": 0,
						"min": 0,
						"name": "lot",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_2037580119",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_35rweyv14r` + "`" + ` ON ` + "`" + `properties` + "`" + ` (\n  ` + "`" + `association` + "`" + `,\n  ` + "`" + `address` + "`" + `,\n  ` + "`" + `unit` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\"",
				"name": "properties",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\" &&\n// Freeze the association\n@request.body.association:changed = false\n",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\""
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\" &&\n@request.body.property.association = @request.body.association &&\n@request.body.user:isset = false",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": true,
						"collectionId": "pbc_2037580119",
						"help": "",
						"hidden": false,
						"id": "relation2347900126",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "property",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation2375276105",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "user",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1579384326",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"exceptDomains": [],
						"help": "",
						"hidden": false,
						"id": "email3885137012",
						"name": "email",
						"onlyDomains": [],
						"presentable": false,
						"required": false,
						"system": false,
						"type": "email"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1146066909",
						"max": 0,
						"min": 0,
						"name": "phone",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select2363381545",
						"maxSelect": 0,
						"name": "type",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"owner",
							"tenant"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool24734798",
						"name": "is_primary",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"help": "",
						"hidden": false,
						"id": "date2502384312",
						"max": "",
						"min": "",
						"name": "start_date",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"help": "null means current",
						"hidden": false,
						"id": "date2220669758",
						"max": "",
						"min": "",
						"name": "end_date",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_2522582908",
				"indexes": [],
				"listRule": "// Logged in\n@request.auth.id != \"\" &&\n// Staff block\n((\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")\n)\n||\n// Resident Block\n(\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.association ?= association &&\n@collection.residents:mine.user ?= @request.auth.id &&\n@collection.residents:mine.property ?= property &&\n@collection.residents:mine.end_date = \"\"\n))",
				"name": "residents",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\" &&\n@request.body.user:changed = false &&\n@request.body.association:changed = false &&\n@request.body.property:changed = false\n",
				"viewRule": "// Logged in\n@request.auth.id != \"\" &&\n// Staff block\n((\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")\n)\n||\n// Resident Block\n(\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.association ?= association &&\n@collection.residents:mine.user ?= @request.auth.id &&\n@collection.residents:mine.property ?= property &&\n@collection.residents:mine.end_date = \"\"\n))"
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\"",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1579384326",
						"max": 0,
						"min": 0,
						"name": "name",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "number4216875057",
						"max": null,
						"min": 0,
						"name": "amount_cents",
						"onlyInt": true,
						"presentable": false,
						"required": true,
						"system": false,
						"type": "number"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select645904403",
						"maxSelect": 0,
						"name": "frequency",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"monthly",
							"quarterly",
							"annual",
							"once"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "date2502384312",
						"max": "",
						"min": "",
						"name": "start_date",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "date"
					},
					{
						"help": "",
						"hidden": false,
						"id": "date2220669758",
						"max": "",
						"min": "",
						"name": "end_date",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool1260321794",
						"name": "active",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3999789178",
				"indexes": [],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")",
				"name": "assessment_schedules",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.memberships:mine.role ?= \"admin\" &&\n@request.body.association:changed = false",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_2037580119",
						"help": "",
						"hidden": false,
						"id": "relation2347900126",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "property",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select2363381545",
						"maxSelect": 0,
						"name": "type",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"charge",
							"payment",
							"credit",
							"late_fee",
							"adjustment",
							"reversal"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "number4216875057",
						"max": null,
						"min": null,
						"name": "amount_cents",
						"onlyInt": true,
						"presentable": false,
						"required": true,
						"system": false,
						"type": "number"
					},
					{
						"help": "",
						"hidden": false,
						"id": "date2555063006",
						"max": "",
						"min": "",
						"name": "posted_on",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "date"
					},
					{
						"help": "",
						"hidden": false,
						"id": "date820172507",
						"max": "",
						"min": "",
						"name": "due_on",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text1843675174",
						"max": 0,
						"min": 0,
						"name": "description",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3999789178",
						"help": "",
						"hidden": false,
						"id": "relation1513624059",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "schedule",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "e.g. 2026-12, used for idempotency",
						"hidden": false,
						"id": "text3317178062",
						"max": 0,
						"min": 0,
						"name": "period",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3289777763",
						"help": "",
						"hidden": false,
						"id": "relation2557951325",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "reverses",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_631030571",
						"help": "",
						"hidden": false,
						"id": "relation1831371789",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "payment",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3289777763",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_xpn97uw2hx` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `schedule` + "`" + `,\n  ` + "`" + `period` + "`" + `,\n  ` + "`" + `property` + "`" + `\n) WHERE 'schedule' != '';",
					"CREATE INDEX ` + "`" + `idx_acvbprphw5` + "`" + ` ON ` + "`" + `ledger_entries` + "`" + ` (\n  ` + "`" + `property` + "`" + `,\n  ` + "`" + `posted_on` + "`" + `\n)"
				],
				"listRule": "@request.auth.id != \"\" &&\n((@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")) ||\n(@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.residents:mine2.user ?= @request.auth.id &&\n@collection.residents:mine2.property ?= property &&\n@collection.residents:mine2.end_date ?= \"\"\n))",
				"name": "ledger_entries",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != \"\" &&\n((@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")) ||\n(@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.residents:mine2.user ?= @request.auth.id &&\n@collection.residents:mine2.property ?= property &&\n@collection.residents:mine2.end_date ?= \"\"\n))"
			},
			{
				"createRule": null,
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_2037580119",
						"help": "",
						"hidden": false,
						"id": "relation2347900126",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "property",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation1103846297",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "payer",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"help": "",
						"hidden": false,
						"id": "number4216875057",
						"max": null,
						"min": null,
						"name": "amount_cents",
						"onlyInt": true,
						"presentable": false,
						"required": true,
						"system": false,
						"type": "number"
					},
					{
						"help": "",
						"hidden": false,
						"id": "number1733422029",
						"max": null,
						"min": null,
						"name": "fee_cents",
						"onlyInt": true,
						"presentable": false,
						"required": false,
						"system": false,
						"type": "number"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select1582905952",
						"maxSelect": 0,
						"name": "method",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"card",
							"ach",
							"check",
							"cash"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "select2063623452",
						"maxSelect": 0,
						"name": "status",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"pending",
							"succeeded",
							"failed",
							"refunded"
						]
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text924342971",
						"max": 0,
						"min": 0,
						"name": "stripe_payment_intent",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text749745336",
						"max": 0,
						"min": 0,
						"name": "stripe_event_id",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"autogeneratePattern": "",
						"help": "check number, memo",
						"hidden": false,
						"id": "text2929936659",
						"max": 0,
						"min": 0,
						"name": "reference",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": false,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "date1833926553",
						"max": "",
						"min": "",
						"name": "received_at",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "date"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_631030571",
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_6xdr3ocyqu` + "`" + ` ON ` + "`" + `payments` + "`" + ` (` + "`" + `stripe_event_id` + "`" + `) WHERE 'stripe_event_id' != '';"
				],
				"listRule": "@request.auth.id != \"\" &&\n((@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")) ||\n(@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.residents:mine2.user ?= @request.auth.id &&\n@collection.residents:mine2.property ?= property &&\n@collection.residents:mine2.end_date ?= \"\"\n))",
				"name": "payments",
				"system": false,
				"type": "base",
				"updateRule": null,
				"viewRule": "@request.auth.id != \"\" &&\n((@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")) ||\n(@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" &&\n@collection.residents:mine2.user ?= @request.auth.id &&\n@collection.residents:mine2.property ?= property &&\n@collection.residents:mine2.end_date ?= \"\"\n))"
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@request.body.author = @request.auth.id &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" && (\n@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\")\n",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation3182418120",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "author",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text724990059",
						"max": 0,
						"min": 0,
						"name": "title",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"convertURLs": false,
						"help": "",
						"hidden": false,
						"id": "editor3685223346",
						"maxSize": 0,
						"name": "body",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "editor"
					},
					{
						"help": "",
						"hidden": false,
						"id": "bool3844597223",
						"name": "pinned",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "bool"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3866499052",
				"indexes": [],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\"",
				"name": "announcements",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\nauthor = @request.auth.id &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n  @collection.memberships:mine.role ?= \"admin\" ||\n  @collection.memberships:mine.role ?= \"board\"\n) &&\n@request.body.association:changed = false &&\n@request.body.author:changed = false",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\""
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\") &&\n@request.body.uploaded_by = @request.auth.id",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": false,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"autogeneratePattern": "",
						"help": "",
						"hidden": false,
						"id": "text724990059",
						"max": 0,
						"min": 0,
						"name": "title",
						"pattern": "",
						"presentable": false,
						"primaryKey": false,
						"required": true,
						"system": false,
						"type": "text"
					},
					{
						"help": "",
						"hidden": false,
						"id": "file2359244304",
						"maxSelect": 0,
						"maxSize": 0,
						"mimeTypes": [
							"application/pdf",
							"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
							"application/vnd.oasis.opendocument.text",
							"application/vnd.oasis.opendocument.text-template",
							"application/msword",
							"text/plain"
						],
						"name": "file",
						"presentable": false,
						"protected": true,
						"required": true,
						"system": false,
						"thumbs": [],
						"type": "file"
					},
					{
						"help": "",
						"hidden": false,
						"id": "select105650625",
						"maxSelect": 0,
						"name": "category",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"bylaws",
							"minutes",
							"financial",
							"forms",
							"other"
						]
					},
					{
						"help": "",
						"hidden": false,
						"id": "select1368277760",
						"maxSelect": 0,
						"name": "visibility",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "select",
						"values": [
							"board",
							"all"
						]
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation3823579430",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "uploaded_by",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3332084752",
				"indexes": [],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n  visibility = \"all\" ||\n  @collection.memberships:mine.role ?= \"admin\" ||\n  @collection.memberships:mine.role ?= \"board\"\n)",
				"name": "documents",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n(@collection.memberships:mine.role ?= \"admin\" || @collection.memberships:mine.role ?= \"board\") &&\n@collection.memberships:mine.status ?= \"active\" &&\n@request.body.association:changed = false &&\n@request.body.uploaded_by:changed = false",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n  visibility = \"all\" ||\n  @collection.memberships:mine.role ?= \"admin\" ||\n  @collection.memberships:mine.role ?= \"board\"\n)"
			},
			{
				"createRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= @request.body.association &&\n@collection.memberships:mine.status ?= \"active\" && (\n  @collection.memberships:mine.role ?= \"admin\" ||\n  @collection.memberships:mine.role ?= \"board\"\n) &&\n@request.body.author = @request.auth.id &&\n@request.body.property.association = @request.body.association\n",
				"deleteRule": null,
				"fields": [
					{
						"autogeneratePattern": "[a-z0-9]{15}",
						"help": "",
						"hidden": false,
						"id": "text3208210256",
						"max": 15,
						"min": 15,
						"name": "id",
						"pattern": "^[a-z0-9]+$",
						"presentable": false,
						"primaryKey": true,
						"required": true,
						"system": true,
						"type": "text"
					},
					{
						"cascadeDelete": true,
						"collectionId": "pbc_2037580119",
						"help": "",
						"hidden": false,
						"id": "relation2347900126",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "property",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"convertURLs": false,
						"help": "",
						"hidden": false,
						"id": "editor3485334036",
						"maxSize": 0,
						"name": "note",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "editor"
					},
					{
						"cascadeDelete": false,
						"collectionId": "_pb_users_auth_",
						"help": "",
						"hidden": false,
						"id": "relation3182418120",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "author",
						"presentable": false,
						"required": false,
						"system": false,
						"type": "relation"
					},
					{
						"cascadeDelete": true,
						"collectionId": "pbc_3736699063",
						"help": "",
						"hidden": false,
						"id": "relation4253360588",
						"maxSelect": 0,
						"minSelect": 0,
						"name": "association",
						"presentable": false,
						"required": true,
						"system": false,
						"type": "relation"
					},
					{
						"hidden": false,
						"id": "autodate2990389176",
						"name": "created",
						"onCreate": true,
						"onUpdate": false,
						"presentable": false,
						"system": false,
						"type": "autodate"
					},
					{
						"hidden": false,
						"id": "autodate3332085495",
						"name": "updated",
						"onCreate": true,
						"onUpdate": true,
						"presentable": false,
						"system": false,
						"type": "autodate"
					}
				],
				"id": "pbc_3772977686",
				"indexes": [],
				"listRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n@collection.memberships:mine.role ?= \"admin\" ||\n@collection.memberships:mine.role ?= \"board\")",
				"name": "property_notes",
				"system": false,
				"type": "base",
				"updateRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n  @collection.memberships:mine.role ?= \"admin\" ||\n  (@collection.memberships:mine.role ?= \"board\" && author = @request.auth.id)\n) &&\n@request.body.author:changed = false &&\n@request.body.property:changed = false &&\n@request.body.association:changed = false",
				"viewRule": "@request.auth.id != \"\" &&\n@collection.memberships:mine.user ?= @request.auth.id &&\n@collection.memberships:mine.association ?= association &&\n@collection.memberships:mine.status ?= \"active\" && (\n@collection.memberships:mine.role ?= \"admin\" ||\n@collection.memberships:mine.role ?= \"board\")"
			}
		]`

		return app.ImportCollectionsByMarshaledJSON([]byte(jsonData), false)
	}, func(app core.App) error {
		return nil
	})
}

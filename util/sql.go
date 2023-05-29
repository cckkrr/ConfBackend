package util

// ChatHistorySqlBeforeMsg 查询历史记录，6个参数，分别是：
// 1. from_user_uuid
// 2. to_entity_uuid
// 3. to_entity_uuid
// 4. from_user_uuid
// 5. 自哪条消息开始 mid
// 6. 查询多少条消息 limit
var ChatHistorySqlBeforeMsg string = `select *
from (
SELECT *
from t_im_message m
WHERE
(
(m.from_user_uuid = ? AND m.to_entity_uuid = ?)
OR (m.from_user_uuid = ? AND m.to_entity_uuid = ?)
)
AND m.created_at < (
SELECT created_at
from t_im_message
where uuid = ?
)
ORDER BY created_at DESC
LIMIT ?
) s
ORDER BY created_at asc`

// ChatHistorySqlBeforeMsgWithoutAsterisk ChatHistorySqlBeforeMid 查询历史记录，6个参数，分别是：
// 1. from_user_uuid
// 2. to_entity_uuid
// 3. to_entity_uuid
// 4. from_user_uuid
// 5. 自哪条消息开始 mid
// 6. 查询多少条消息 limit
var ChatHistorySqlBeforeMsgWithoutAsterisk string = `select *
from (
SELECT m.id, m.uuid, m.msg_type, m.text_type_text, m.file_type_uri, m.from_user_uuid, m.to_entity_uuid, m.created_at
from t_im_message m
WHERE
(
(m.from_user_uuid = ? AND m.to_entity_uuid = ?)
OR (m.from_user_uuid = ? AND m.to_entity_uuid = ?)
)
AND m.created_at < (
SELECT created_at
from t_im_message
where uuid = ?
)
ORDER BY created_at DESC
LIMIT ?
) s
ORDER BY created_at asc`

// ChatHistorySqlLatest 查询历史记录，5个参数，分别是：
// 1. from_user_uuid
// 2. to_entity_uuid
// 3. to_entity_uuid
// 4. from_user_uuid
// 5. 查询多少条消息 limit
var ChatHistorySqlLatest string = `select *
from (
SELECT *
from t_im_message m
WHERE
(
(m.from_user_uuid = ? AND m.to_entity_uuid = ?)
OR (m.from_user_uuid = ? AND m.to_entity_uuid = ?)
)
ORDER BY created_at DESC
LIMIT ?
) s
ORDER BY created_at asc`

// ChatHistorySqlLatestWithoutAsterisk ChatHistorySqlLatest 查询历史记录，5个参数，分别是：
// 1. from_user_uuid
// 2. to_entity_uuid
// 3. to_entity_uuid
// 4. from_user_uuid
// 5. 查询多少条消息 limit
var ChatHistorySqlLatestWithoutAsterisk string = `select *
from (
SELECT m.id, m.uuid, m.msg_type, m.text_type_text, m.file_type_uri, m.from_user_uuid, m.to_entity_uuid, m.created_at
from t_im_message m
WHERE
(
(m.from_user_uuid = ? AND m.to_entity_uuid = ?)
OR (m.from_user_uuid = ? AND m.to_entity_uuid = ?)
)
ORDER BY created_at DESC
LIMIT ?
) s
ORDER BY created_at asc`

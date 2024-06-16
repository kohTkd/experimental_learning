table "accounts" {
  schema  = schema.api_dev
  collate = "utf8mb4_bin"
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "email" {
    null = false
    type = varchar(255)
  }
  column "password" {
    null = false
    type = varchar(255)
  }
  column "created_at" {
    null = false
    type = timestamp
  }
  column "update_at" {
    null = false
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
  index "email" {
    unique  = true
    columns = [column.email]
  }
}
table "atlas_schema_revisions" {
  schema  = schema.api_dev
  collate = "utf8mb4_bin"
  column "version" {
    null = false
    type = varchar(255)
  }
  column "description" {
    null = false
    type = varchar(255)
  }
  column "type" {
    null     = false
    type     = bigint
    default  = 2
    unsigned = true
  }
  column "applied" {
    null    = false
    type    = bigint
    default = 0
  }
  column "total" {
    null    = false
    type    = bigint
    default = 0
  }
  column "executed_at" {
    null = false
    type = timestamp
  }
  column "execution_time" {
    null = false
    type = bigint
  }
  column "error" {
    null = true
    type = longtext
  }
  column "error_stmt" {
    null = true
    type = longtext
  }
  column "hash" {
    null = false
    type = varchar(255)
  }
  column "partial_hashes" {
    null = true
    type = json
  }
  column "operator_version" {
    null = false
    type = varchar(255)
  }
  primary_key {
    columns = [column.version]
  }
}
schema "api_dev" {
  charset = "utf8mb4"
  collate = "utf8mb4_unicode_ci"
}

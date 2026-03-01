# hashes

Hashes of the file's binary content, if available. Read-only.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `crc32Hash` | string | No | The CRC32 value of the file (if available). Read-only. |
| `quickXorHash` | string | No | A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only. |
| `sha1Hash` | string | No | SHA1 hash for the contents of the file (if available). Read-only. |
| `sha256Hash` | string | No | SHA256 hash for the contents of the file (if available). Read-only. |


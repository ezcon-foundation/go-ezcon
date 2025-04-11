## Easy Connect with Blockchain

Golang execution layer implementation of the Ezcon protocol.


```text
go-ezcon/
├── cmd/                    // Lệnh thực thi
│   ├── ezcon/              // Lệnh chính (tương tự geth)
│   │   └── main.go         // Điểm vào chương trình
│   └── utils/              // Công cụ phụ (ví dụ: tạo khóa)
├── common/                 // Tiện ích chung
│   ├── hexutil/            // Chuyển đổi hex
│   ├── math/               // Toán học (uint64, big.Int)
│   └── types/              // Kiểu dữ liệu cơ bản
├── core/                   // Logic cốt lõi
│   ├── ledger/             // Quản lý ledger (tương tự blockchain)
│   │   ├── ledger.go       // Định nghĩa Ledger và xử lý trạng thái
│   │   └── validator.go    // Xử lý giao dịch bởi validator
│   ├── state/              // Trạng thái tài khoản (RocksDB)
│   └── types/              // Định nghĩa Account, Transaction
├── crypto/                 // Mã hóa và chữ ký
│   ├── ecdsa.go            // Chữ ký ECDSA
│   └── hash.go             // Hash SHA-256
├── consensus/              // Đồng thuận RPCA
│   ├── rpca/               // Triển khai RPCA
│   │   └── rpca.go         // Logic đồng thuận
├── kyc/                    // Xử lý KYC
│   ├── provider.go         // KYC provider logic
│   └── types.go            // KYCData và SignatureResult
├── network/                // Giao tiếp P2P
│   ├── p2p.go              // Mạng P2P với libp2p
│   └── websocket.go        // WebSocket cho client-validator
├── rpc/                    // API RPC
│   └── server.go           // JSON-RPC server
├── db/                     // Cơ sở dữ liệu
│   ├── rocksdb/            // Trạng thái ledger (RocksDB)
│   └── boltdb/             // Lịch sử ledger (BoltDB)
├── trie/                   // Cấu trúc dữ liệu (nếu cần Merkle Tree)
├── vendor/                 // Thư viện bên thứ ba
├── go.mod                  // Module Golang
└── README.md               // Tài liệu
```
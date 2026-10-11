## CoinFlow

기존 [FinFlow 프로젝트](https://github.com/zzzyoonnn/finflow-backend)에서는 중앙화된 금융 시스템에서 **거래의 정합성과 멱등성, 장애 복구**를 중심으로 구현했습니다.

CoinFlow에서는 이러한 문제들이 **분산 환경에서는 어떻게 해결되는지** 직접 구현하며 비교하고자 합니다.

### 중앙화 금융 시스템과 블록체인 비교

| FinFlow에서 경험한 것       | CoinFlow에서 구현할 것       |
| --------------------------- | ---------------------------- |
| Transfer                    | Transaction                  |
| DB Transaction              | State Transition             |
| Idempotency Key             | Transaction ID / Nonce       |
| 중복 이체 방지              | Double Spending 방지         |
| DB Lock                     | 트랜잭션 / 블록 검증         |
| 거래 내역                   | Blockchain Ledger            |
| Kafka Event                 | P2P Transaction / Block 전파 |
| Transactional Outbox 정합성 | 노드 간 데이터 동기화        |
| 장애 복구                   | Chain / State Sync           |

FinFlow가 하나의 신뢰할 수 있는 서버와 데이터베이스를 중심으로
거래의 정합성을 보장했다면,

CoinFlow에서는 여러 노드가 동일한 거래와 상태를 공유하는 환경에서
**어떻게 하나의 일관된 상태를 만들어 가는지**를 직접 구현합니다.

---

## Core Components

CoinFlow에서는 다음 요소들을 중심으로 블록체인의 기본 구조를 구현합니다.

- Block
- Transaction
- UTXO
- Wallet
- Digital Signature
- Mempool
- Mining / Proof of Work
- P2P Network
- Node
- Block Validation

using Go = import "/go.capnp";

@0xa1e8437c93150c11;

$Go.package("main");
$Go.import("dapp/main");

struct AdvanceRequest @0x9f15675dca75dd4d {
    union {
        add :group {
            operand @0 :Int64;
        }
        mul :group {
            operand @1 :Int64;
        }
        div :group {
            operand @2 :Int64;
        }
    }
}

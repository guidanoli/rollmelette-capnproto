using Go = import "/go.capnp";

@0xa1e8437c93150c11;

$Go.package("main");
$Go.import("dapp/main");

struct AdvanceRequest @0x9f15675dca75dd4d {
    union {
        add @0 : Int64;
        mul @1 : Int64;
        div @2 : Int64;
    }
}

struct InspectRequest @0xdeef9407cb2e1380 {
    union {
        value @0 : Void;
        opCount @1 : Void;
    }
}

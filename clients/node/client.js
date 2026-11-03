
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const path = require("path");

const PROTO_PATH = path.resolve(__dirname, "../../user.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const proto = grpc.loadPackageDefinition(packageDefinition).user;

const client = new proto.UserService(
    "localhost:50051",
    grpc.credentials.createInsecure()
);

client.GetUser({ id: 1 }, (error, response) => {
    if (error) {
        console.error("Gagal mengambil data pengguna:", error.message);
        client.close();
        return;
    }

    console.log("Data pengguna berhasil diterima:");
    console.log(`ID    : ${response.id}`);
    console.log(`Nama  : ${response.name}`);
    console.log(`Email : ${response.email}`);

    client.close();
});

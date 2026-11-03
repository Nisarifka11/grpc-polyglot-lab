
import grpc
import user_pb2
import user_pb2_grpc


def run():
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)

        request = user_pb2.UserRequest(id=1)
        response = stub.GetUser(request)

        print("Data pengguna berhasil diterima:")
        print(f"ID    : {response.id}")
        print(f"Nama  : {response.name}")
        print(f"Email : {response.email}")


if __name__ == "__main__":
    run()
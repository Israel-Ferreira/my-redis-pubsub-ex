import json
from redis import StrictRedis
from models.car import convert_dict_to_product



if __name__ == "__main__":
    print("Subscriber Redis -  Pub/Sub")

    redis_client = StrictRedis("localhost", 6379)

    subscriber = redis_client.pubsub()
    subscriber.subscribe("cars")


    while True:
        message =  subscriber.get_message()

        if message:
            if type(message["data"]) != int:
                json_msg = json.loads(message["data"], object_hook=convert_dict_to_product)
                print(json_msg.model)

            


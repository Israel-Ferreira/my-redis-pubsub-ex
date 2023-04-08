class Car:
    def __init__(self, uuid,  model, manufacturer, description, license_number,  price):
        self.uuid = uuid
        self.model =  model
        self.manufacturer =  manufacturer
        self.description = description
        self.license_number =  license_number
        self.price =  price



def convert_dict_to_product(json_dict: dict):
    return Car(
        json_dict["uid"], 
        json_dict["model"], 
        json_dict["manufacturer"],
        json_dict["description"],
        json_dict["license_number"],
        json_dict["price"]
    )

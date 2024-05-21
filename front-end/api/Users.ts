
import type { ApiResponseList, LoginBody, RegisterBody } from "~/types";
import Abstract from "./abstract";
import { Actions } from "~/constants/actions";

export class Users extends Abstract {

    static async getContacts() {
        const res= await this.get("/contacts", Actions.GetUsers);
        
        return res.list;
    }
    

}








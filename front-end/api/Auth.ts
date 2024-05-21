import type { LoginBody, RegisterBody } from "~/types";
import Abstract from "./abstract";
import { Actions } from "~/constants/actions";

export class Auth extends Abstract {

    static login(body: LoginBody) {
        return this.post("/users/login", body, Actions.Login);
    }
    static register(body: RegisterBody) {
        return this.post("/users/register", body, Actions.Register);
    }
}


import type { ApiResponseList, ChatMessageBody, LoginBody, RegisterBody } from "~/types";
import Abstract from "./abstract";
import { Actions } from "~/constants/actions";

export class Message extends Abstract {

    static async getChatMessages(query: ChatMessageBody) {
        const res = await this.get("/messages", Actions.GetChatMessages, { user_id: query.userId, receiver_id: query.receiverId });

        return res.list;
    }


}








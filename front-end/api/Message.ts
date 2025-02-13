
import type { ApiResponseList, ChatMessageBody, LoginBody, MarkAsReadBody, RegisterBody } from "~/types";
import Abstract from "./abstract";
import { Actions } from "~/constants/actions";

export class Message extends Abstract {

    static async getChatMessages(query: ChatMessageBody) {
        const res = await this.get("/messages", Actions.GetChatMessages, { user_id: query.userId, receiver_id: query.receiverId });

        return res.list;
    }

    static async markAsRead(data: MarkAsReadBody) {
        const res = await this.post("/messages/mark-as-read", {messageId: data.messageId, readAllLatest:data.readAllLatest,receiverId: data.receiverId }, Actions.MarkAsRead);

        return res;
    }
}








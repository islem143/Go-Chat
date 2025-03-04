export interface ApiResponseList {
  list: any[],
}

export interface RegisterBody {
  name: string,
  email: string,
  password: string
}
export interface LoginBody {

  email: string,
  password: string
}
export interface User {
  id: string
  name: string,
  email: string,

}

export interface ChatMessageBody {

  receiverId: string | null,
  userId: string | null
}

export interface MarkAsReadBody {
  messageId: string | null,
  receiverId: string,
  readAllLatest: boolean

}

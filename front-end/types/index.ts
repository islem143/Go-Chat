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
export interface User{
  id:number
  name: string,
  email: string,
  
}
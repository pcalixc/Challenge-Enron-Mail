export interface IEmail {
  message_id: string;
  date: string;
  from: string;
  to: string;
  subject: string;
  content: string;
}

export interface IHit{
	_index:  string 
  _type: string,
    _id:    string 
    _source: IEmail    
}

export interface ISearchReq {
  type:string,
  from: string,
  to: string,
}

export interface IServerErrorResponse {
  errorStatus: boolean,
  errorCode: number,
  errorMessage: string,
}

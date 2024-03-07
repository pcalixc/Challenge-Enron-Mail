export interface IEmail {
  message_id: string;
  date: string;
  from: string;
  to: string;
  subject: string;
  content: string;
}


export interface Hit{
	_index:  string 
  _type: string,
    _id:    string 
    _source: IEmail    
}

export interface SearchReq {
  type:string,
  from: string,
  to: string,
}

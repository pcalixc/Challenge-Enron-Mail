export interface Email {
  message_id: string;
  date: string;
  from: string;
  to: string;
  subject: string;
  cc: string;
  mime_version: string;
  content_type: string;
  content_transfer_encoding: string;
  bcc: string;
  x_from: string;
  x_to: string;
  x_cc: string;
  x_bcc: string;
  x_folder: string;
  x_origin: string;
  x_file_name: string;
  content: string;
}


export interface Hit{
	_index:  string 
  _type: string,
    _id:    string 
    _source: Email    
}


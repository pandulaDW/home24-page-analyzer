export interface RequestBody {
  url: string;
}

export interface ResponseBody {
  html_version: string;
  title: string;
  heading_count: {
    h1_count: number;
    h2_count: number;
    h3_count: number;
    h4_count: number;
    h5_count: number;
    h6_count: number;
  };
  link_count: {
    internal_link_count: number;
    external_link_count: number;
    inaccessible_link_count: number;
  };
  is_login_form: boolean;
}

export interface ErrorResponseBody {
  err_message: string;
  err_description: string;
  err_status_code: number;
}

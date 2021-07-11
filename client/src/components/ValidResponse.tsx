import React from "react";
import { ResponseBody } from "../models/urlModels";

interface Props {
  urlResponse: ResponseBody;
}

const ValidResponse: React.FC<Props> = ({ urlResponse }) => {
  return (
    <div className="valid-response">
      <ul>
        <li>
          <span>HTML Version - </span>
          <span>{urlResponse.html_version}</span>
        </li>
        <li>
          <span>Title - </span>
          <span>{urlResponse.title}</span>
        </li>
        <li>
          <span>Header Information</span>
          <ul className="valid-response_sub-ul">
            <li>
              <span>h1 Count - </span>
              <span>{urlResponse.heading_count.h1_count}</span>
            </li>
            <li>
              <span>h2 Count - </span>
              <span>{urlResponse.heading_count.h2_count}</span>
            </li>
            <li>
              <span>h3 Count - </span>
              <span>{urlResponse.heading_count.h3_count}</span>
            </li>
            <li>
              <span>h4 Count - </span>
              <span>{urlResponse.heading_count.h4_count}</span>
            </li>
            <li>
              <span>h5 Count - </span>
              <span>{urlResponse.heading_count.h5_count}</span>
            </li>
            <li>
              <span>h6 Count - </span>
              <span>{urlResponse.heading_count.h6_count}</span>
            </li>
          </ul>
        </li>
        <li>
          <span>Link Information</span>
          <ul className="valid-response_sub-ul">
            <li>
              <span>Internal Link Count - </span>
              <span>{urlResponse.link_count.internal_link_count}</span>
            </li>
            <li>
              <span>External Link Count - </span>
              <span>{urlResponse.link_count.external_link_count}</span>
            </li>
            <li>
              <span>Inaccessible Link Count - </span>
              <span>{urlResponse.link_count.inaccessible_link_count}</span>
            </li>
          </ul>
        </li>
        <li>
          <span>Form Included - </span>
          <span>{urlResponse.is_login_form ? "true" : "false"}</span>
        </li>
      </ul>
    </div>
  );
};

export default ValidResponse;

import React from "react";
import "./../css/url_div.css";
import DenseTable from "./Table";

const tempData = [
  {
    url: "https://www.instagram.com/",
    short_url: "kushagra_wa",
    Hits: 0,
    Comments: "Whatsapp referral",
    Medium: "Whatsapp",
    Source: "Whatsapp",
    Campaign: "Test-Campaign-1",
    Keyword: "Table",
    Created_at: 0,
  },
  {
    url: "https://www.instagram.com/",
    short_url: "kushagra_w1",
    Hits: 1,
    Comments: "Whatsapp referral",
    Medium: "Whatsapp",
    Source: "Whatsapp",
    Campaign: "Test-Campaign-1",
    Keyword: "Table",
    Created_at: 0,
  },
];
export default function Dashboard(props) {
  return (
    // <div className="Url_tab_container">
    //   <div className="Url_tab">
    //     <div className="Short">Shortened URLs</div>
    //     <div className="Long">Original</div>
    //     <div className="Hits" style={{ color: "white" }}>
    //       Hits
    //     </div>
    //   </div>
    //   {[...props.urls].map((u) => (
    //     <div className="Url_tab">
    //       <div className="Short">
    //         <a href={`${BASE}` + u.short_url}>{`${BASE}` + u.short_url}</a>
    //       </div>
    //       <div className="Long">{u.url}</div>
    //       <div className="Hits">{u.Hits}</div>
    //     </div>
    //   ))}
    <div
      style={{
        width: "80%",
        margin: "auto",
      }}
    >
      <DenseTable rows={props.urls} />
    </div>
    // </div>
  );
}

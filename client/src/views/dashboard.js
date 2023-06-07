import React from 'react'
import './../css/url_div.css'
export default function Dashboard(props) {
  const BASE = 
  // process.env.BASE_URL
  'https://go-tinly.onrender.com'
  +'/r/'
  return (

    <div className="Url_tab_container">
      <div className="Url_tab"><div className="Short">Shortened URLs</div><div className="Long">Original</div><div className="Hits" style={{color:"white"}}>Hits</div></div>
      {
        [...props.urls].map(u => <div className="Url_tab"><div className="Short"><a href={`${BASE}`+ u.short_url}>{`${BASE}`+ u.short_url}</a></div><div className="Long">{u.url}</div><div className="Hits">{u.Hits}</div></div>)
      }
    </div>

  )
}

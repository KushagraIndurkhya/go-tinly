import React from 'react'
import { Eclipse } from "react-loading-io";
import "./../css/component.css"

export default function Loading() {
        return <div className="loader"><Eclipse size={64}  /></div>;
}

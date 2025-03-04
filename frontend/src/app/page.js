'use client'

import { getExcel } from "@/req/spesification";
import styles from "./page.module.css";
import { useState } from "react";

export default function Home() {

  const [text, setText] = useState("");

  const handleDownload = async () => {
    console.log("Button Pressed");

    const blob = await getExcel(text);
    if (!blob) return;


    const currentDate = new Date().toISOString().split("T")[0];

    const url = window.URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `Spesifikasyon_${currentDate}.xlsx`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    window.URL.revokeObjectURL(url);
  };

  return (
    <div>
      <div className={styles.page}>
        <main className={styles.main}>
          <h1>Enter Text</h1>
          <input
            type="text"
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="Type something..."
            style={{ padding: "10px", fontSize: "16px" }}
          />
          <p>Typed: {text}</p>
          <button onClick={handleDownload}>
            Download Excel
          </button>
        </main >
      </div >
    </div>

  );
}

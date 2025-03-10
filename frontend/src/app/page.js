'use client'

import { getDev, getExcel, getOEMTable, getTableData } from "@/req/spesification";
import styles from "./page.module.css";
import { useState } from "react";
import { DataTable } from "@/components/SpecTable";
import { OEMTable } from "@/components/OEMTable";

export default function Home() {
  const [data, setData] = useState([]);
  const [oemdata, setOemData] = useState([]);
  const [text, setText] = useState("");

  const handleDownload = async () => {
    console.log("Request received...");

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

  const handleTable = async () => {
    const tableData = await getTableData(text);
    setData(tableData)
  }

  const handleOEMTAble = async () => {
    const oemdata = await getOEMTable(text);
    setOemData(oemdata)
  }

  const handleDev = async () => {
    const tabel = await getDev(text);
  }

  return (
    <div>
      <div className={styles.page}>
        <main className={styles.main}>
          <h1>Supap Kodu Girin:</h1>
          <input
            type="text"
            value={text}
            onChange={(e) => setText(e.target.value)}
            placeholder="Supap Kodları..."
            style={{ width: "500px", padding: "10px", fontSize: "16px" }}
          />
          <div>
            <p style={{ width: "320px" }}>Yazılan kodlar: {text}</p>
            <button onClick={handleOEMTAble}>
              OEM no dan bul
            </button>
            <button onClick={handleTable}>
              Sonucu Göster
            </button>
            <button onClick={handleDownload}>
              Exceli İndir
            </button>
          </div>
          {data.length > 0 ? <DataTable data={data} /> : <p></p>}
          {oemdata.length > 0 ? <OEMTable data={oemdata} /> : <p></p>}
        </main >
      </div >
    </div >

  );
}


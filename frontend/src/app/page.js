'use client'

import { getExcel, getTableData } from "@/req/spesification";
import styles from "./page.module.css";
import { useState } from "react";

export default function Home() {
  const [data, setData] = useState([]);
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
            <button onClick={handleTable}>
              Sonucu Göster
            </button>
            <button onClick={handleDownload}>
              Exceli İndir
            </button>
          </div>
          {data.length > 0 ? <DataTable data={data} /> : <p></p>}
        </main >
      </div >
    </div >

  );
}

const DataTable = ({ data }) => {
  return (
    <div style={{ margin: '20px', textAlign: 'center' }}>
      <h2>Spesifikasyon</h2>
      <table className={styles.tablestyle}>
        <thead style={{ justifyContent: 'space-between' }}>
          <tr className={styles.table}>
            <th>Supar Kodu</th>
            <th>Model</th>
            <th>IN/EX</th>
            <th>OEM</th>
          </tr>
        </thead>
        <tbody>
          {data.map((user) => (
            <tr className={styles.table} key={user.suparCode}>
              <td >{user.suparCode}</td>
              <td >{user.makeModel}</td>
              <td >{user.type}</td>
              <td >{user.originalCode}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

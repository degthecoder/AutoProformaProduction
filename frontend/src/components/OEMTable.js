import styles from "../app/page.module.css";


export const  OEMTable = ({ data }) => {
    return (
      <div style={{ margin: '20px', textAlign: 'center' }}>
        <h2>Spesifikasyon</h2>
        <table className={styles.tablestyle}>
          <thead style={{ justifyContent: 'space-between' }}>
            <tr className={styles.table}>
              <th>Supar Kodu</th>
              <th>OEM</th>
              <th>Grup Kodu</th>
            </tr>
          </thead>
          <tbody>
            {data.map((user) => (
              <tr className={styles.table} key={user.suparCode}>
                <td >{user.suparCode}</td>
                <td >{user.originalCode}</td>
                <td >{user.urunGrupKodu}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };
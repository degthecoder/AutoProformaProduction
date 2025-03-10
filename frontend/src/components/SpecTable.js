import styles from "../app/page.module.css";

export const  DataTable = ({ data }) => {
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
              <th>UGK</th>
            </tr>
          </thead>
          <tbody>
            {data.map((product) => (
              <tr className={styles.table} key={product.suparCode}>
                <td >{product.suparCode}</td>
                <td >{product.makeModel}</td>
                <td >{product.type}</td>
                <td >{product.originalCode}</td>
                <td> {product.urunGrupKodu}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };
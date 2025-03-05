'use server'

export async function getExcel(text) {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;

    const response = await fetch(`${apiUrl}/getExcel`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: text,
    });

    if (!response.ok) {
        throw new Error("Failed to fetch Excel file");
    }

    console.log("End of function.")
    return await response.blob();
}

export async function getTableData(text) {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    console.log(apiUrl);
    const response = await fetch(`${apiUrl}/getTable`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: text,
    }).catch(error=> console.error(error));

    console.log("HERE", response.body)
    if (!response.ok) {
        throw new Error("Failed to fetch table");
    }

    return await response.json();
}


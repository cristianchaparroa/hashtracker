
export async function fetchTransactions(
    address:string
): Promise<any> {
    console.log("--> querying transactions: ", address);
    const baseUrl = 'https://hashai-na7w9.ondigitalocean.app';
    const url = `${baseUrl}/transactions/${address}?level=1`;
    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
            },
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const message = await response.json();
        const data =  JSON.stringify(message, null, 2);
        return JSON.parse(data);
    } catch (error) {
        throw new Error(`Failed to fetch wallet transactions: ${error}`);
    }
}

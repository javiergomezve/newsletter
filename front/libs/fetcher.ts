import axios from 'axios'

const fetcher = async <T>(url: string): Promise<T> => {
    try {
        const response = await axios.get(url);
        return response.data;
    } catch (error) {
        throw new Error('Error fetching data');
    }
};


export default fetcher

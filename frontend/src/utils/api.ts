const API_BASE_URL = import.meta.env.VITE_API_URL;

export const apiFetch = async <T>(endpoint: string): Promise<T | null> => {
    try {
        const response = await fetch(`${API_BASE_URL}${endpoint}`);
        if (!response.ok) {
            throw new Error(`${response.status} ${response.statusText}`);
        }
        return await response.json();
    } catch (err) {
        console.error("Error fetching data:", err);
        return null;
    }
};
export interface Stock {
    ticker: string;
    company: string;
    price: number;
    change: number;
    action: string;
    target_from: number;
    target_to: number;
    rating_to: string;
}
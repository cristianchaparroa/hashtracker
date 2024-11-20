import {XMTPContext, SkillResponse, getUserInfo} from "@xmtp/message-kit";
import {fetchTransactions} from "../services/transactions.js";
import { shortenUrl } from 'shaveurl';
import {createReport} from "../services/reports.js";

export async function transactionsHandler(context: XMTPContext): Promise<SkillResponse | undefined> {
    const {
        message: {
            content: {
                params: { address },
            },
        },
    } = context;

    const transactions = await fetchTransactions(address);
    const url = `https://hashtracker.vercel.app/?hash=${transactions.mermaid}`
    const shavedURL = await shortenUrl(url,"isgd");
    return {
        code: 200,
        message: `The following diagram illustrates the ${address} address behaviour. ${shavedURL}`
    };
}

export async function reportHandler(context: XMTPContext): Promise<SkillResponse | undefined> {
    const {
        message: {
            content: {
                params: { address },
            },
        },
    } = context;

    const result = await createReport(address);
    console.log("---> transaction result:", result);
    return { code: 200, message: `${address} reported...` };
}
import { createClient } from '@urql/core';

export const clientFactory = () => {
    return createClient({
        url: 'http://localhost:3000/graphql',
        exchanges: [/** ... */],
    });
};
import { defineConfig } from 'orval';

export default defineConfig({
    taskApi: {
        input: '../task-api/docs/swagger.yaml',
        output: {
            target: './src/generated/task-api.ts',
            httpClient: 'axios',
            namingConvention: 'camelCase',
            override: {
                mutator: {
                    path: './src/lib/axios.ts',
                    name: 'axiosInstance'
                }
            }
        },
        hooks: {
            afterAllFilesWrite: 'npx prettier --write'
        }
    }
});

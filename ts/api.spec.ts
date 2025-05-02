import { test, expect, request } from "@playwright/test";

test.describe("API Health", () => {
  test("GET /health returns OK", async () => {
    const apiAddr = process.env.API_ADDR || "http://127.0.0.1:8080";
    const api = await request.newContext({ baseURL: apiAddr });
    const resp = await api.get("/health");
    expect(resp.status()).toBe(200);
    const body = await resp.json();
    expect(body.status).toBe("OK");
  });
});

import { test, expect, request } from "@playwright/test";

test.describe("API Integration", () => {
  test("GET /api/v1/resource returns data", async () => {
    const apiAddr = process.env.API_ADDR || "http://127.0.0.1:8080";
    const api = await request.newContext({ baseURL: apiAddr });
    const resp = await api.get("/api/v1/resource"); // adjust path
    expect(resp.ok()).toBeTruthy();
    const data = await resp.json();
    expect(Array.isArray(data)).toBeTruthy();
  });
});

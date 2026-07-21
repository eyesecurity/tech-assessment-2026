import { useQuery } from "@tanstack/react-query";
import { apiFetch } from "./client";

export interface HealthResponse {
  status: string;
}

// Example hook: proves the frontend can reach the Go backend.
// Reuse this React Query pattern for your own data fetching.
export function useHealth() {
  return useQuery({
    queryKey: ["health"],
    queryFn: () => apiFetch<HealthResponse>("/health"),
  });
}

import { useHealth } from "./api/useHealth";

export default function App() {
  return (
    <div className="min-h-screen bg-gray-50">
      <header className="border-b border-gray-200 bg-white">
        <div className="mx-auto max-w-7xl px-6">
          <div className="flex h-14 items-center justify-between">
            <span className="font-semibold text-gray-900">Eye Security</span>
            <ConnectivityBadge />
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-7xl px-6 py-8">
        <div className="flex flex-col gap-4">
          <h1 className="text-xl font-semibold text-gray-900">Detections</h1>
          <div className="rounded-lg border border-dashed border-gray-300 bg-white p-12 text-center text-gray-500">
            Build the detections triage view here.
          </div>
        </div>
      </main>
    </div>
  );
}

// Example component: calls GET /api/health to show the backend is reachable.
function ConnectivityBadge() {
  const { data, isLoading, error } = useHealth();

  const { label, className } = error
    ? { label: "backend: unreachable", className: "bg-red-100 text-red-700" }
    : isLoading
      ? { label: "backend: checking…", className: "bg-gray-100 text-gray-500" }
      : {
          label: `backend: ${data?.status ?? "unknown"}`,
          className: "bg-green-100 text-green-700",
        };

  return (
    <span className={`rounded px-2 py-0.5 text-xs font-medium ${className}`}>
      {label}
    </span>
  );
}

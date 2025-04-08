// Follow this setup guide to integrate the Deno language server with your editor:
// https://deno.land/manual/getting_started/setup_your_environment
// This enables autocomplete, go to definition, etc.

// Setup type definitions for built-in Powerbase Runtime APIs
import "jsr:@skorpland/functions-js/edge-runtime.d.ts"

console.log("Hello from Functions!")

Deno.serve(async (req) => {
  const { name } = await req.json()
  const data = {
    message: `Hello ${name}!`,
  }

  return new Response(
    JSON.stringify(data),
    { headers: { "Content-Type": "application/json" } },
  )
})

/* To invoke locally:

  1. Run `powerbase start` (see: https://powerbase.club/docs/reference/cli/powerbase-start)
  2. Make an HTTP request:

  curl -i --location --request POST '{{ .URL }}' \
    --header 'Authorization: Bearer {{ .Token }}' \
    --header 'Content-Type: application/json' \
    --data '{"name":"Functions"}'

*/

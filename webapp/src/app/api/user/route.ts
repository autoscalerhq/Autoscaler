import { ensureSuperTokensInit } from "~/config/supertokens/backend";
import { NextResponse, NextRequest } from "next/server";
import { withSession } from "supertokens-node/nextjs";

ensureSuperTokensInit();

export function GET(request: NextRequest) {
  return withSession(request, async (err, session) => {
    if (err) {
      return NextResponse.json(err, { status: 500 });
    }
    if (!session) {
      return new NextResponse("Authentication required", { status: 401 });
    }

    return NextResponse.json({
      note: "Fetch any data from your application for authenticated user after using verifySession middleware",
      userId: session.getUserId(),
      sessionHandle: session.getHandle(),
      // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
      accessTokenPayload: session.getAccessTokenPayload(),
    });
  });
}

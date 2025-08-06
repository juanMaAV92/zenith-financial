import { NextRequest, NextResponse } from 'next/server';

const BACKEND_URL = process.env.BACKEND_URL || 'http://localhost:8080/zenith-financial';

export async function GET(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  return handleRequest(request, params.path, 'GET');
}

export async function POST(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  return handleRequest(request, params.path, 'POST');
}

export async function PUT(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  return handleRequest(request, params.path, 'PUT');
}

export async function DELETE(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  return handleRequest(request, params.path, 'DELETE');
}

export async function PATCH(
  request: NextRequest,
  { params }: { params: { path: string[] } }
) {
  return handleRequest(request, params.path, 'PATCH');
}

async function handleRequest(
  request: NextRequest,
  path: string[],
  method: string
) {
  try {
    const url = `${BACKEND_URL}/${path.join('/')}`;
    
    // Preparar headers
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
    };

    // Obtener body para métodos que lo requieren
    let body: string | undefined;
    if (['POST', 'PUT', 'PATCH'].includes(method)) {
      const requestBody = await request.text();
      if (requestBody) {
        body = requestBody;
      }
    }

    // Realizar la petición al backend
    const response = await fetch(url, {
      method,
      headers,
      body,
    });

    // Obtener la respuesta
    const responseData = await response.text();
    
    // Devolver la respuesta con el mismo status
    return new NextResponse(responseData, {
      status: response.status,
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error) {
    console.error('Proxy error:', error);
    return NextResponse.json(
      { 
        code: 'PROXY_ERROR',
        messages: ['Error de conexión con el servidor'] 
      },
      { status: 500 }
    );
  }
}

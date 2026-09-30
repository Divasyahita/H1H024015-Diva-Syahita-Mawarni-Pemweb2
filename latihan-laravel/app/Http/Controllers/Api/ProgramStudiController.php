<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Resources\MahasiswaResource;
use App\Models\ProgramStudi;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class ProgramStudiController extends Controller
{
    public function mahasiswa(
        Request $request,
        ProgramStudi $programStudi
    ): JsonResponse {
        $perHalaman = min(
            $request->integer('per_halaman', 10),
            100
        );

        $mahasiswa = $programStudi
            ->mahasiswa()
            ->with('programStudi')
            ->paginate($perHalaman);

        return response()->json([
            'sukses' => true,
            'program_studi' => [
                'id' => $programStudi->id,
                'kode' => $programStudi->kode,
                'nama' => $programStudi->nama,
                'jenjang' => $programStudi->jenjang,
            ],
            'data' => MahasiswaResource::collection($mahasiswa),
        ]);
    }
}
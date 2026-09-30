<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Http\Requests\StoreMahasiswaRequest;
use App\Http\Requests\UpdateMahasiswaRequest;
use App\Http\Resources\MahasiswaResource;
use App\Models\Mahasiswa;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class MahasiswaController extends Controller
{
        public function index(Request $request)
    {
        $kueri = Mahasiswa::query()->with('programStudi');

        // Memilih kolom yang ingin ditampilkan
        $fields = $request->query('fields');

        if ($fields) {
            $kolomDiizinkan = [
                'id',
                'nim',
                'nama',
                'email',
                'angkatan',
                'ipk',
                'aktif',
                'program_studi_id',
                'created_at',
            ];

            $fields = array_filter(
                array_map('trim', explode(',', $fields)),
                fn ($field) => in_array($field, $kolomDiizinkan, true)
            );

            // ID dan foreign key tetap diambil agar resource dan relasi tetap aman
            $fields = array_unique([
                'id',
                ...$fields,
                'program_studi_id',
            ]);

            $kueri->select($fields);
        }

        if ($request->filled('cari')) {
            $kataKunci = $request->query('cari');

            $kueri->where(function ($sub) use ($kataKunci) {
                $sub->where('nama', 'like', '%' . $kataKunci . '%')
                    ->orWhere('nim', 'like', '%' . $kataKunci . '%');
            });
        }

        if ($request->filled('angkatan')) {
            $kueri->where('angkatan', $request->integer('angkatan'));
        }

        if ($request->filled('program_studi_id')) {
            $kueri->where(
                'program_studi_id',
                $request->integer('program_studi_id')
            );
        }

        $urutan = $request->query('urut', 'nama');
        $arah = $request->query('arah', 'asc');

        $kolomDiizinkan = ['nama', 'nim', 'angkatan', 'ipk'];

        if (in_array($urutan, $kolomDiizinkan, true)) {
            $kueri->orderBy(
                $urutan,
                $arah === 'desc' ? 'desc' : 'asc'
            );
        }

        $perHalaman = min(
            $request->integer('per_halaman', 10),
            100
        );

        return MahasiswaResource::collection(
            $kueri->paginate($perHalaman)
        );
    }

    public function store(StoreMahasiswaRequest $request): JsonResponse
    {
        $mahasiswa = Mahasiswa::create($request->validated());

        $mahasiswa->load('programStudi');

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil dibuat',
            'data' => new MahasiswaResource($mahasiswa),
        ], 201);
    }

    public function show(Mahasiswa $mahasiswa): JsonResponse
    {
        $mahasiswa->load('programStudi');

        return response()->json([
            'sukses' => true,
            'data' => new MahasiswaResource($mahasiswa),
        ]);
    }

    public function update(
        UpdateMahasiswaRequest $request,
        Mahasiswa $mahasiswa
    ): JsonResponse {
        $mahasiswa->update($request->validated());

        $mahasiswa->load('programStudi');

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil diperbarui',
            'data' => new MahasiswaResource($mahasiswa),
        ]);
    }

    public function destroy(Mahasiswa $mahasiswa): JsonResponse
    {
        $mahasiswa->delete();

        return response()->json([
            'sukses' => true,
            'pesan' => 'Data mahasiswa berhasil dihapus',
        ]);
    }
}
<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class MahasiswaResource extends JsonResource
{
    public function toArray(Request $request): array
    {
        $fields = $request->query('fields');

        // Kalau fields tidak dikirim, tampilkan semua data seperti biasa
        if (!$fields) {
            return [
                'id' => $this->id,
                'nim' => $this->nim,
                'nama' => $this->nama,
                'email' => $this->email,
                'angkatan' => $this->angkatan,
                'ipk' => (float) $this->ipk,
                'aktif' => $this->aktif,

                'program_studi' => $this->whenLoaded('programStudi', function () {
                    return [
                        'id' => $this->programStudi->id,
                        'kode' => $this->programStudi->kode,
                        'nama' => $this->programStudi->nama,
                    ];
                }),

                'dibuat_pada' => $this->created_at->toIso8601String(),
            ];
        }

        $fields = array_map('trim', explode(',', $fields));

        $data = [];

        if (in_array('id', $fields)) {
            $data['id'] = $this->id;
        }

        if (in_array('nim', $fields)) {
            $data['nim'] = $this->nim;
        }

        if (in_array('nama', $fields)) {
            $data['nama'] = $this->nama;
        }

        if (in_array('email', $fields)) {
            $data['email'] = $this->email;
        }

        if (in_array('angkatan', $fields)) {
            $data['angkatan'] = $this->angkatan;
        }

        if (in_array('ipk', $fields)) {
            $data['ipk'] = (float) $this->ipk;
        }

        if (in_array('aktif', $fields)) {
            $data['aktif'] = $this->aktif;
        }

        if (in_array('dibuat_pada', $fields)) {
            $data['dibuat_pada'] = $this->created_at->toIso8601String();
        }

        if (in_array('program_studi', $fields) && $this->relationLoaded('programStudi')) {
            $data['program_studi'] = [
                'id' => $this->programStudi->id,
                'kode' => $this->programStudi->kode,
                'nama' => $this->programStudi->nama,
            ];
        }

        return $data;
    }
}
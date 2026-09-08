<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class MatakuliahController extends Controller
{
    public function index(Request $request)
    {
        $daftarMatakuliah = [
            [
                'kode' => 'TK101',
                'nama' => 'Pemrograman Web II',
                'sks' => 3,
            ],
            [
                'kode' => 'TK102',
                'nama' => 'Sistem Operasi',
                'sks' => 3,
            ],
            [
                'kode' => 'TK103',
                'nama' => 'Jaringan Komputer',
                'sks' => 3,
            ],
            [
                'kode' => 'TK104',
                'nama' => 'Sistem Kendali',
                'sks' => 2,
            ],
            [
                'kode' => 'TK105',
                'nama' => 'Internet of Things',
                'sks' => 2,
            ],
        ];

        $kataKunci = $request->query('q', '');

        if ($kataKunci !== '') {
            $daftarMatakuliah = array_filter(
                $daftarMatakuliah,
                function ($matakuliah) use ($kataKunci) {
                    return stripos($matakuliah['nama'], $kataKunci) !== false;
                }
            );
        }

        return view('matakuliah.index', [
            'daftarMatakuliah' => $daftarMatakuliah,
            'kataKunci' => $kataKunci
        ]);
    }

    public function show(string $kode)
    {
        return view('matakuliah.show', [
            'kode' => $kode
        ]);
    }
}
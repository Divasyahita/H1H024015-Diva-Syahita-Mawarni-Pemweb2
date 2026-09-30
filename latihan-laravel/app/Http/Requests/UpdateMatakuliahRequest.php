<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Validation\Rule;

class UpdateMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        return [
            'kode' => [
                'sometimes',
                'required',
                'string',
                'max:20',
                Rule::unique('matakuliahs', 'kode')
                    ->ignore($this->route('matakuliah')),
            ],
            'nama' => ['sometimes', 'required', 'string', 'max:255'],
            'sks' => ['sometimes', 'required', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'required', 'integer', 'min:1', 'max:8'],
        ];
    }
}
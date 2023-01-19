(vim.api.nvim_create_namespace :debug)
(vim.api.nvim_buf_clear_namespace 0 41 0 -1)

(vim.api.nvim_buf_set_extmark 0 41 1 1
                              {:id 4
                               :virt_text_pos :overlay
                               :virt_lines [[["Hi there is this cool? Making a long string hereereeeeeeeeeeeeeeeeeeeeeeeeeee hidden messages should also be visible otherwise this is pretty useless"
                                              :DiffText]]]})
